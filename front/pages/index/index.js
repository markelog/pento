import React, { useState, useEffect } from "react";

import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";

import { get as getUser } from "../../lib/user";

import Layout from "../../components/layout";
import StartStop from "../../components/start-stop";
import Calendar from "../../components/calendar";

import { setStatus } from "./tracker";
import { getTracks } from "./tracks";

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

function Index({ user, status }) {
  const classes = useStyles();

  const [tracks, setTracks] = useState([]);

  const requestTracks = async () => {
    const data = await getTracks(user.email);
    setTracks(data);
  };

  useEffect(() => {
    requestTracks();
  }, [setTracks]);

  function update(active, name) {
    if (active === false) {
      requestTracks();
    }
    setStatus(user.email, active, name);
  }

  return (
    <>
      <Link href="/api/logout" color="inherit" className={classes.link}>
        logout
      </Link>
      <Layout user={user}>
        <StartStop email={user.email} active={status.active} onClick={update} />
      </Layout>
      <Calendar email={user.email} events={tracks} />
    </>
  );
}

Index.getInitialProps = async ({ req, res }) => {
  const data = await getUser(req);

  // Redirect to login if user is not there
  if (data === null) {
    res.writeHead(302, {
      Location: "/api/login"
    });
    res.end();
    return;
  }

  return data;
};

export default Index;
