require 'json'

RESULT_FILE_NAME = "all-apis.openapi.json"
OPENAPI_DIR = File.join(File.dirname(__FILE__), "..", "data/docs/openapi/")
RESULT_FILE = File.join(OPENAPI_DIR, RESULT_FILE_NAME)

openapi_files = Dir["#{OPENAPI_DIR}/*.json"].reject {|f| File.basename(f) == RESULT_FILE_NAME }

output = JSON.parse(File.read(openapi_files.shift))
output["info"]["title"] = RESULT_FILE

openapi_files.each do |openapi_file|
  openapi_json = JSON.parse(File.read(openapi_file))

  output['paths'].merge!(openapi_json['paths'])
  output['components']['schemas'].merge!(openapi_json['components']['schemas'])
end

File.write(RESULT_FILE, JSON.dump(output))

puts "Successfully combined files in #{OPENAPI_DIR} into #{RESULT_FILE}."
