import axios, { AxiosProgressEvent } from "axios";
import { CONSTANTS } from "src/constants";

export const getAllDestinations = () => {
  return axios.get(`${CONSTANTS.MAIN_SERVER}/destination/all`);
};

export const getDestinationByName = (name: string) => {
  return axios.get(`${CONSTANTS.MAIN_SERVER}/destination/${name}`);
};

export const createDestination = (name: string, url: string) => {
  return axios({
    url: `${CONSTANTS.MAIN_SERVER}/destination/create`,
    method: "POST",
    data: {
      name,
      url,
    },
  });
};

export const updateDestination = (id: number, name: string, url: string) => {
  return axios({
    url: `${CONSTANTS.MAIN_SERVER}/destination/update`,
    method: "PUT",
    data: {
      id,
      name,
      url,
    },
  });
};

export const deleteDestination = (name: string) => {
  return axios.delete(`${CONSTANTS.MAIN_SERVER}/destination/delete/${name}`, {
    method: "DELETE",
  });
};
