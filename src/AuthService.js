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
    if (response.status != 200) {
      throw response.status;
    }
    return response.text();
  };
  signUp = async (user) => {
    const url = config.backendUrl;

    var formData = new FormData();

    for (let key in user) {
      formData.append(key, user[key]);
    }

    const options = {
      method: "POST",
      credentials: 'include', // include, *same-origin, omit
      body: user,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default AuthService;
