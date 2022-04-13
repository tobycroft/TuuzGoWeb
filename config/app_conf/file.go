package app_conf

const FileSavePath = "public"
const FilePathCreateByDay = true  //create day's privilige is higher than date, if this on the following Date will be no longer avail
const FilePathCreateByDate = true //if you want to save the file into current month's folder, make sure you have turn this on and deactivated the CreateByDay setting
const FileNameSecurity = true     //this will turn original file name into MD5 to avoid path attack
