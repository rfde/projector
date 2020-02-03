function sendPostRequest(caller, endpoint, parameter) {
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			if (!caller.classList.contains("button--outlined")) {
				caller.classList.remove("button--filled");
				caller.classList.add("button--outlined");
				caller.blur();
			}
		}
	};

	var data = new FormData();
	data.append('parameter', parameter);

	xhttp.open("POST", "/ajax/" + endpoint, true);
	// xhttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
	xhttp.send(data);
}