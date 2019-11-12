import { useState } from "react";

import { makeStyles } from "@material-ui/core/styles";

import Button from "@material-ui/core/Button";
import Input from "@material-ui/core/Input";

const useStyles = makeStyles(theme => ({
  button: {
    margin: "10px auto",
    display: "block"
  },
  input: {
    display: "none"
  }
}));

export default ({ email, active, onClick }) => {
  const classes = useStyles();
  const [status, setStatus] = useState(!!active);
  const [name, setName] = useState("");

  return (
    <>
      <Input
        autoFocus={true}
        placeholder="What's sup? Tag me"
        color="secondary"
        value={name}
        fullWidth={true}
        onChange={e => setName(e.target.value)}
        disabled={status}
      />
      <Button
        variant="outlined"
        color={status ? "secondary" : "default"}
        className={classes.button}
        onClick={() => {
          onClick(!status, name);
          setStatus(!status);
          setName("");
        }}
      >
        {status ? "Stop" : "Start"}
      </Button>
    </>
  );
};
