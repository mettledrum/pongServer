(function() {
	var app = angular.module('pongApp', []);

	app.controller('PicController', ['$http', function ($http) {
		var picServer = this;
		picServer.pictures = [];
		picServer.errors = [];

		var allPicsEndpoint = 'http://localhost:8080/pong-pics/';

		// parse pic file names from server
		$http.get(allPicsEndpoint).
			success(function(data) {
				var parser = new DOMParser();
				var doc = parser.parseFromString(data, 'text/html');
				var picLinks = doc.firstChild.querySelectorAll('a');

				for (var i = 0; i < picLinks.length; ++i) {
					var picHash = {};
					picHash['image'] = allPicsEndpoint + picLinks[i].innerHTML;
					// TODO: make picture file name epoch in millis
					millis = new Date().valueOf() + i * 3699123;
					picHash['time'] = millis;

					picServer.pictures.push(picHash);
				}
			}).
			error(function() {
				picServer.errors.push("Darn... could not retrieve pictures from pong server.")
			});
	}]);
})();