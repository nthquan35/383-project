var submit = document.getElementById('submit');
var form = document.getElementById('form');
// var ouput = document.getElementById('output');
var tbody = document.querySelector("tbody");

submit.addEventListener("click", sendRequest);

function sendRequest(event){
	event.preventDefault();
	// var element = {
	// 	input : input.value
	// }
	var xhr = new XMLHttpRequest();
	// console.log(form.method);
    xhr.open(form.method, form.action, true);
   	xhr.setRequestHeader('Content-Type', 'application/json');
   	xhr.send();
   	// console.log(JSON.stringify(element));
   	xhr.onreadystatechange = function() { 
	    if (xhr.readyState == 4 && xhr.status == 200){
			var response = JSON.parse(xhr.responseText).array;
			// console.log(response.length);
			for (i = 0; i < response.length ; i++){
				array = response[i];
				let tr = document.createElement("tr");
				for (j = 0; j < array.length; j++){
					let td = document.createElement("td");
					td.appendChild(document.createTextNode(array[j]));
					tr.appendChild(td);
					if (array.length == 1) {
						let td1 = document.createElement("td");
						td1.appendChild(document.createTextNode("This is a prime."));
						tr.appendChild(td1);
					}
				}
				tbody.appendChild(tr);
			}


			// if (response === "Success"){
			// 	successMessage.innerHTML = "<span style='color: green;'>Your Rectangle has been added!</span>";
			// 	fade(successMessage);
			// 	setTimeout(location.reload.bind(location), 2000);
			// }
			// else{
			// 	successMessage.innerHTML = "<span style='color: red;'>An error occurs!</span>";				
			// 	fade(successMessage);
			// 	setTimeout(location.reload.bind(location), 2000);
			// }	
	    }
    };
	// else{
	// 	successMessage.innerHTML = "<span style='color: red;'>Missing values.</span>";
	// 	fade(successMessage);
	// 	setTimeout(location.reload.bind(location), 2000);
	// }
}