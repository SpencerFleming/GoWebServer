$(document).ready(function(){
	$("#funbutton").mouseover(function(){
		var stringx = Math.floor(Math.random() * 500) + 50 + 'px';
		var stringy = Math.floor(Math.random() * 500) + 50 + 'px';
		$("#fundiv button").css('margin-left', stringx);
		$("#fundiv button").css('margin-top', stringy);
	});
});
