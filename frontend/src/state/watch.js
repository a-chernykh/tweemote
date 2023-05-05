export const watch = (getState, path) => {
  let curV = getState()[path];

  return (cb) => {
    return () => {
      let newV = getState()[path];
      if (newV !== curV) {
        curV = newV;
        cb();
      }
    };
  }
};
