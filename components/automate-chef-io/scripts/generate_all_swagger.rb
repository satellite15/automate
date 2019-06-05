require 'json'

RESULT_FILE_NAME = "all-apis.swagger.json"
SWAGGER_DIR = File.join(File.dirname(__FILE__), "..", "data/docs/api_chef_automate/")
RESULT_FILE = File.join(SWAGGER_DIR, RESULT_FILE_NAME)
CONTENT_KEYS = ["definitions", "paths"]

swagger_files = Dir["#{SWAGGER_DIR}/*.json"].reject {|f| File.basename(f) == RESULT_FILE_NAME }

output = JSON.parse(File.read(swagger_files.shift))
output["info"]["title"] = RESULT_FILE

swagger_files.each do |swagger_file|
  swagger_json = JSON.parse(File.read(swagger_file))

  CONTENT_KEYS.each do |key|
    output[key].merge!(swagger_json[key])
  end
end

File.write(RESULT_FILE, JSON.dump(output))
