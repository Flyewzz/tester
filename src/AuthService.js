import config from "./config";

class AuthService {
  login = async (login, password) => {
    const url = config.backendUrl + '/auth/login';

    var formData = new FormData();

    var params = {
      login: login,
      password: password,
    };

    for (let i in params) {
      formData.append(i, params[i]);
    }

    const options = {
      method: "POST",
      body: formData,
    };

    const request = new Request(url, options);
    const response = await fetch(request);

    return response.text().then(text => {
      return {
        text,
        status: response.status,
      }
    });
  }

  signUp = async (user) => {
    const url = config.backendUrl + '/auth/signup';

    var formData = new FormData();

    for (let key in user) {
      formData.append(key, user[key]);
    }

    const options = {
      method: "POST",
      body: formData,
    };

    const request = new Request(url, options);
    const response = await fetch(request);

    return response.text().then(text => {
      return {
        text,
        status: response.status,
      }
    });
  };
}

export default AuthService;
