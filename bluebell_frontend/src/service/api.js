
import axios from 'axios';
axios.defaults.baseURL = "api";
axios.interceptors.request.use( (config) => {
  const token = sessionStorage.getItem('token')
    if (token ) { 
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