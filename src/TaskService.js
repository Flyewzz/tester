import config from "./config";

class TaskService {
  getInfo = async (id) => {
    const url = config.backendUrl + `?id=${id}`;
    const options = {
      method: "GET",
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default TaskService;
