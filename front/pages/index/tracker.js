import fetch from "isomorphic-unfetch";
import moment from "moment";

const API = process.env.API;

function now() {
  return moment(new Date()).format("YYYY-MM-DDTHH:mm:ssZ");
}

function send(body) {
  return fetch(`${API}/tracks`, {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(body)
  });
}

export function setStatus(email, active, name) {
  if (active) {
    return send({
      email,
      name,
      active,
      start: now()
    });
  }

  return send({
    email,
    name,
    active,
    stop: now()
  });
}
