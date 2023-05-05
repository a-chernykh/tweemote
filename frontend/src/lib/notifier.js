class Notifier {
  // https://github.com/igorprado/react-notification-system
  setAdapter(adapter) {
    this.adapter = adapter;
  }

  error(msg, title='Error', opts={}) {
    let n = Object.assign({
      title: title,
      message: msg,
      level: 'error',
      autoDismiss: 0
    }, opts);
    this.adapter.addNotification(n);
  }

  info(msg) {
    let n = {
      message: msg,
      level: 'info',
      autoDismiss: 3
    };
    this.adapter.addNotification(n);
  }
}

export let notifier = new Notifier();
