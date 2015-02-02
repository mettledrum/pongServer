(function() {
	var app = angular.module('pongApp', ['ngSanitize']);

	app.controller('PicController', ['$http', function ($http) {
		var picServer = this;
		picServer.pictures = [];
		picServer.errors = [];

		var allPicsEndpoint = 'http://localhost:8080/pong-pics/';
		var gifizeEndpoint = 'http://localhost:8080/gifize';

		picServer.gifPics = [];
		picServer.gifPicNames = [];
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
					picServer.pictures.push(picHash);
				}
			}).
			error(function() {
				picServer.errors.push("Could not retrieve pictures from pong server.")
			});

		picServer.addFile = function(imageName) {
			// id is needed to make pictures' html strings unique
			picHtml = '<img id="' + gifPicIdx.toString() + '" class="small-image" src="' + imageName + '"/>'
			picServer.gifPics.push(picHtml);
			picServer.gifPicNames.push(imageName);
			gifPicIdx++;
		};

		picServer.clearFiles = function() {
			picServer.gifPics = [];
			picServer.gifPicNames = [];
			gifPicIdx = 0;
		};

		picServer.gifize = function() {
			$http.post(gifizeEndpoint, {'names': picServer.gifPicNames}).
				success(function() {
					console.log('worked');
				}).
				error(function() {
					console.log('did not work');
				});
		};
	}]);
})();