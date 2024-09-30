import moment from "moment";

export const formatDateTime = (date: Date): string => {
  return moment(date).format("MMMM Do YYYY, h:mm:ss a");
};

export const getHourFormat = (date: Date): string => {
  return moment(date).format("LT");
};

export const getDateFormat = (date: Date): string => {
  return moment(date).format("MMM Do YY");
};
