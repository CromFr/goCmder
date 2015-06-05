function sendSource(){
	var request = new XMLHttpRequest();
	
	var formData = new FormData();
	formData.append('source', document.getElementById('source').value);
	formData.append('cmd', document.getElementById('console-cmd').value);
	
	request.open('POST', '/run', true);
	request.send(formData);
}

window.onload = function() {
	if(!location){
		alert("Your browser is shitty, get a new one");
		return
	}
	
	var address = "ws://"+location.hostname+":"+location.port+"/ws";
	var sock = new WebSocket(address);
	sock.onopen = function() {
		console.log("connected to "+address);
	}
	sock.onclose = function(e) {
		console.log("connection closed (" + e.code + ")");
	}
	sock.onmessage = function(e) {
		var container = document.getElementById('console-content');
		container.innerHTML = e.data.replace(/\n/g, "<br/>");
	}
};