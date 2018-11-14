var Client = require('node-rest-client').Client;

var client = new Client();

// direct way
client.get("http://148.100.98.30:3000/api/org.acme.sample.Sensor/1", function (data, response) {
	// parsed response body as js object
  console.log('first method way')
	console.log(data);
	// raw response
	//console.log(response);
});

// registering remote methods
client.registerMethod("jsonMethod", "http://148.100.98.30:3000/api/org.acme.sample.Sensor/1", "GET");

client.methods.jsonMethod(function (data, response) {
	// parsed response body as js object
  console.log('second method way')
	console.log(data);
	// raw response
	//console.log(response);
});
