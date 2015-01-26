(function() {
	var app = angular.module('pongApp', ['ngSanitize']);

	app.controller('PicController', ['$http', function ($http) {
		var picServer = this;
		picServer.pictures = [];
		picServer.errors = [];

		var allPicsEndpoint = 'http://localhost:8080/pong-pics/';
		var gifizeEndpoint = 'http://localhost:8080/gifize';

		picServer.gifPics = [];
		picServer.gifPicTimes = [];
		var gifPicIdx = 0;

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

		picServer.addFile = function(imageName, imageTime) {
			// id is needed to make pictures' html strings unique
			picHtml = '<img id="' + gifPicIdx.toString() + '" class="small-image" src="' + imageName + '"/>'
			picServer.gifPics.push(picHtml);
			picServer.gifPicTimes.push(imageTime);
			gifPicIdx++;
		};

		picServer.clearFiles = function() {
			picServer.gifPics = [];
			picServer.gifPicTimes = [];
			gifPicIdx = 0;
		};

		picServer.gifize = function() {
			$http.post(gifizeEndpoint, {msg: 'ahoy mundo!'}).
				success(function() {
					console.log('worked');
				}).
				error(function() {
					console.log('did not work');
				});
		};
	}]);
})();