import config from "./config";

class TaskService {
  getInfo = async (id, token) => {
    const url = config.backendUrl + `?id=${id}`;
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

export default TaskService;
