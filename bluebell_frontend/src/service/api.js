
import axios from 'axios';
axios.defaults.baseURL = "http://127.0.0.1:8080/api/v1/";
axios.interceptors.request.use((config) => {
  const token = localStorage.getItem('bluebellToken')
    if (token) { 
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
}, (error) => {
	return Promise.reject(error);
});

axios.interceptors.response.use(
	response => {
		if (response.status === 200) {
			return Promise.resolve(response.data);
		} else {
			return Promise.reject(response.data);
		}
	},
	(error) => {
		console.log('error', error);
	}
);

export default axios;