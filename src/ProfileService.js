import config from "./config";

class ProfileService {

  getUserTasks = async(token) => {
    const url = config.backendUrl + `/profile/attempts`;
    const options = {
      method: "GET",
      headers: new Headers({
        "Authorization": `Bearer ${token}`,
      }),
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    if (response.status !== 200) {
      throw response.status;
    }
    return response.json();
  };
}

export default ProfileService;
