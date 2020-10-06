import config from "./config";

class VerdictService {
  getVerdicts = async (code) => {
    const url = config.backendUrl + "/test";
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
