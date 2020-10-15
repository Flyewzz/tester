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

import AuthService from "./AuthService";
import { useState } from "react";
import Cookie from "js-cookie";
import { Redirect } from "react-router-dom";
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
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function LoginPage() {
  const classes = useStyles();
  let authService = new AuthService();

  const [state, setState] = useState({
    isLogged: Cookie.get("token"),
    login: "",
    password: "",
    message: {
      opened: false,
      text: "",
    },
  });

  // const [cookies, setCookie] = useCookies('token');

  const handleChange = (e) => {
    const value = e.target.value;
    setState({
      ...state,
      [e.target.name]: value,
    });
  };

  const onSubmitClick = (e) => {
    e.preventDefault();
    authService
      .login(state.login, state.password)
      .then((result) => {
        if (result.status !== 200) {
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
          case 401:
            err.text =
              "Произошла ошибка авторизации. Неправильный логин или пароль.";
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
          },
        });
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

  if (state.isLogged) {
    return <Redirect to="/1"></Redirect>;
  }
  return (
    <>
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
              Вход
            </Typography>
            <form className={classes.form} noValidate>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="login"
                label="Электронная почта / Логин"
                name="login"
                autoComplete="email"
                value={state.login}
                onChange={handleChange}
                autoFocus
              />
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Пароль"
                type="password"
                id="password"
                autoComplete="current-password"
                onChange={handleChange}
                value={state.password}
              />
              <FormControlLabel
                control={<Checkbox value="remember" color="primary" />}
                label="Запомнить меня"
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
                onClick={onSubmitClick}
              >
                Войти
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link href="#" variant="body2">
                    Забыли пароль?
                  </Link>
                </Grid>
                <Grid item>
                  <Link href="/signup" variant="body2">
                    {"Зарегистрироваться"}
                  </Link>
                </Grid>
              </Grid>
            </form>
          </div>
          <Box mt={8}>
            <Copyright />
          </Box>
        </Container>
      </MuiThemeProvider>
    </>
  );
}
