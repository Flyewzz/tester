import React from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

import Theme from "./Theme";
import { MuiThemeProvider } from "@material-ui/core/styles";

import Cookie from "js-cookie";
import { Redirect } from "react-router-dom";
import { useState } from "react";
import AuthService from "./AuthService";
import { Snackbar } from "@material-ui/core";
import Alert from "@material-ui/lab/Alert";

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Repetory
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function SignUp() {
  const classes = useStyles();
  const [state, setState] = useState({
    isLogged: Cookie.get("token"),
    name: "",
    login: "",
    email: "",
    password: "",
    message: {
      opened: false,
      text: '',
    }
  });

  const handleChange = (e) => {
    const value = e.target.value;
    setState({
      ...state,
      [e.target.name]: value,
    });
  };

  const handleCloseMessage = (event, reason) => {
    setState({
      ...state,
      message: {
        opened: false,
      },
    });
  };

  let authService = new AuthService();

  const onSubmitClick = (e) => {
    e.preventDefault();
    let user = {
      name: state.name,
      login: state.login,
      email: state.email,
      password: state.password,
    };
    authService
      .signUp(user)
      .then((result) => {
        if (result.status !== 201) {
          throw result;
        }
        Cookie.set("token", result.text);
        setState({
          ...state,
          isLogged: true,
        });
      })
      .catch((err) => {
        switch (err.status) {
          case 400:
            break;
          case 500:
            err.text = "Произошла внутренняя ошибка сервера.";
            break;
          default:
            err.text = "Произошла неизвестная ошибка.";
            break;
        }
        setState({
          ...state,
          message: {
            opened: true,
            text: err.text,
          }
        });
      });
  };

  if (state.isLogged) {
    return <Redirect to="/1"></Redirect>;
  }

  return (
    <MuiThemeProvider theme={Theme}>
      <Container component="main" maxWidth="xs">
      <Snackbar
          anchorOrigin={{
            vertical: "top",
            horizontal: "center",
          }}
          // autoHideDuration={8000}
          open={state.message.opened}
          onClose={handleCloseMessage}
        >
          <Alert severity="error">{state.message.text}</Alert>
        </Snackbar>
        <CssBaseline />
        <div className={classes.paper}>
          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Регистрация
          </Typography>
          <form className={classes.form} noValidate>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  autoComplete="fname"
                  name="name"
                  variant="outlined"
                  required
                  fullWidth
                  id="name"
                  label="Имя и фамилия"
                  value={state.name}
                  onChange={handleChange}
                  autoFocus
                />
              </Grid>
              {/* <Grid item xs={12} sm={6}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="lastName"
                  label="Last Name"
                  name="lastName"
                  autoComplete="lname"
                />
              </Grid> */}
              <Grid item xs={12}>
                <TextField
                  autoComplete="fname"
                  name="login"
                  variant="outlined"
                  required
                  fullWidth
                  id="login"
                  label="Логин"
                  value={state.login}
                  onChange={handleChange}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="email"
                  label="Электронная почта"
                  value={state.email}
                  onChange={handleChange}
                  name="email"
                  autoComplete="email"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  name="password"
                  label="Пароль"
                  value={state.password}
                  onChange={handleChange}
                  type="password"
                  id="password"
                  autoComplete="current-password"
                />
              </Grid>
              {/* <Grid item xs={12}>
                <FormControlLabel
                  control={
                    <Checkbox value="allowExtraEmails" color="primary" />
                  }
                  label="I want to receive inspiration, marketing promotions and updates via email."
                />
              </Grid> */}
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              onClick={onSubmitClick}
              className={classes.submit}
            >
              Зарегистрироваться
            </Button>
            <Grid container justify="flex-end">
              <Grid item>
                <Link href="/login" variant="body2">
                  Уже есть аккаунт?
                </Link>
              </Grid>
            </Grid>
          </form>
        </div>
        <Box mt={5}>
          <Copyright />
        </Box>
      </Container>
    </MuiThemeProvider>
  );
}
