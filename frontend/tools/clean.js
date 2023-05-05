import fs from 'fs-extra';

function clean() {
  return new Promise((resolve, reject) => {
    fs.remove('./dist', err => {
      if (err) {
        reject(err);
      } else {
        resolve('Success');
      }
    });
  });
}

export default clean;
