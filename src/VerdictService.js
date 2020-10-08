import config from "./config";

class VerdictService {
  getVerdicts = async (id, code) => {
    const url = config.backendUrl + `/test/${id}`;
    const formData = new FormData();
    formData.append("code", code);
    const options = {
      method: "POST",
      body: formData,
    };

    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default VerdictService;
