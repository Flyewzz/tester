import config from "../config";

class VerdictService {
  getVerdicts = async (id, code, token) => {
    const url = config.backendUrl + `/test/${id}`;
    const formData = new FormData();
    formData.append("code", code);
    const options = {
      method: "POST",
      body: formData,
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

export default VerdictService;
