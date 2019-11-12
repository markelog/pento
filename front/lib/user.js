import fetch from "isomorphic-unfetch";

import auth0 from "./auth0";

const API = process.env.API;

export async function get(req) {
  const result = await auth0.getSession(req);

  if (result === null) {
    return null;
  }

  const { user } = result;
  const { email } = user;

  const [status, tracks] = await Promise.all([
    getStatus(email),
    getTracks(email)
  ]);

  return {
    user,
    status,
    tracks
  };
}

function getStatus(email) {
  return fetch(`${API}/user/status/${email}`)
    .then(r => r.json())
    .then(data => data.payload);
}

function getTracks(email) {
  return fetch(`${API}/tracks/${email}`)
    .then(r => r.json())
    .then(data => data.payload);
}
