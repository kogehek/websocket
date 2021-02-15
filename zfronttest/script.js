let url = "ws://localhost:8080/ws";

let socket = new WebSocket(url);

socket.onmessage = function (event) {
  let node = document.getElementById("server_messages");
  let selfMassage = document.createElement("div");
  selfMassage.classList.add("string_messages")
  selfMassage.innerHTML = event.data;
  node.insertBefore(selfMassage, node.childNodes[0]);
};

function getUID() {
  socket.send(JSON.stringify({
      Method:"getName"
      
  }));
}

function broadcast() {
  let message = document.getElementById("message_input").value;
  let selfMassage = document.createElement("div");
  selfMassage.innerHTML = message;
  document.getElementById("message_input").value = "";
  socket.send(JSON.stringify({
    Method:"broadcast", 
    Body: message
  }));
}

function registration() {
  let email = document.getElementById("email_input").value;
  let password = document.getElementById("password_input").value;
  socket.send(JSON.stringify({
    Method:"registration",
    Body:{"email":email,"password": password},
  }));
}

function auth() {
  let email = document.getElementById("email_input").value;
  let password = document.getElementById("password_input").value;
  socket.send(JSON.stringify({
    Method:"auth",
    Body:{"email":email,"password": password},
  }));
}