package service

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	secrets "github.com/chef/automate/api/external/secrets"
	cfgmgmt "github.com/chef/automate/api/interservice/cfgmgmt/service"
	"github.com/chef/automate/components/compliance-service/api/reporting"
	"github.com/chef/automate/components/data-feed-service/config"
	"github.com/chef/automate/components/data-feed-service/dao"
	"github.com/chef/automate/lib/cereal"
	"github.com/chef/automate/lib/grpc/secureconn"
)

var (
	dataFeedNotifierTaskName = cereal.NewTaskName("data-feed-notifier")
)

type DataFeedNotifierTask struct {
	cfgMgmt   cfgmgmt.CfgMgmtClient
	secrets   secrets.SecretsServiceClient
	reporting reporting.ReportingServiceClient
	db        *dao.DB
	manager   *cereal.Manager
}

type DataFeedNotifierTaskParams struct {
	DataFeedMessages map[string]datafeedMessage
}

type DataFeedNotifierTaskResults struct {
}

func NewDataFeedNotifierTask(dataFeedConfig *config.DataFeedConfig, connFactory *secureconn.Factory, db *dao.DB, manager *cereal.Manager) (*DataFeedNotifierTask, error) {

	cfgMgmtConn, err := connFactory.Dial("config-mgmt-service", dataFeedConfig.CfgmgmtConfig.Target)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to config-mgmt-service")
	}

	secretsConn, err := connFactory.Dial("secrets-service", dataFeedConfig.SecretsConfig.Target)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to secrets-service")
	}

	complianceConn, err := connFactory.Dial("compliance-service", dataFeedConfig.ComplianceConfig.Target)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to compliance-service")
	}

	return &DataFeedNotifierTask{
		secrets:   secrets.NewSecretsServiceClient(secretsConn),
		reporting: reporting.NewReportingServiceClient(complianceConn),
		cfgMgmt:   cfgmgmt.NewCfgMgmtClient(cfgMgmtConn),
		db:        db,
		manager:   manager,
	}, nil
}

func (d *DataFeedNotifierTask) Run(ctx context.Context, task cereal.Task) (interface{}, error) {

	params := DataFeedWorkflowParams{}
	err := task.GetParameters(&params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse task parameters")
	}
	log.Infof("DataFeedNotifierTask.Run %v", params)
	datafeedMessages := params.NotifierTaskParams.DataFeedMessages
	data := make([]datafeedMessage, 0, len(datafeedMessages))
	for _, value := range datafeedMessages {
		data = append(data, value)
	}

	if len(datafeedMessages) == 0 {
		return nil, nil
	}

	destinations, err := d.db.ListDBDestinations()
	for destination := range destinations {
		log.Debugf("Destination name %v", destinations[destination].Name)
		log.Debugf("Destination url %v", destinations[destination].URL)
		log.Debugf("Destination secret %v", destinations[destination].Secret)

		username, password, err := GetCredentials(ctx, d.secrets, destinations[destination].Secret)

		if err != nil {
			log.Errorf("Error retrieving credentials, cannot send asset notification: %v", err)
			// TODO error handling - need some form of report in automate that indicates when data was sent and if it was successful
		} else {
			// build and send notification for this rule
			notification := datafeedNotification{username: username, password: password, url: destinations[destination].URL, data: data}

			client := NewDataClient()
			err = send(client, notification)
			if err != nil {
				handleSendErr(notification, params.ComplianceTaskParams.FeedStart, params.ComplianceTaskParams.FeedEnd, err)
			}
		}
	}

	return nil, nil
}
