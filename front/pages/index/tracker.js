import fetch from "isomorphic-unfetch";
import moment from "moment";

const API = process.env.API;

function now() {
  return moment(new Date()).format("YYYY-MM-DDTHH:mm:ssZ");
}

function send(body) {
  fetch(`${API}/tracks`, {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(body)
  });
}

export function setStatus(email, status, name) {
  if (status) {
    send({
      email,
      name,
      active: status,
      start: now()
    });

    return;
  }

  send({
    email,
    name,
    active: status,
    stop: now()
  });
}
