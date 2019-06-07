import {UserState} from '../model/state/userState';

export class WS {
  static socket;

  static send(msg: string) {
    try {
      WS.socket.send(msg);
    } catch (e) {
      this.SetupFromScratch();
      setTimeout(() => {
        WS.send(msg);
      }, 1000);
    }
  }

  static handler(x: any) {
    console.log('No handler...');
    console.log(x);
  }

  static SetupFromScratch() {
    WS.Setup('ws://' + window.location.host + '/ws/entity-sync/', UserState.getMyKey());
  }

  static Setup(url: string, key: string) {
    WS.socket = new WebSocket(url);

    WS.socket.onopen = function (e) {
      console.log('[open] Connection established, send -> server');
      const action = {
        action: 'secret',
        body: key,
      };
      WS.socket.send(JSON.stringify(action));
    };

    WS.send = (msg: string) => {
      try {
        WS.socket.send(msg);
      } catch (e) {
        this.SetupFromScratch();
        setTimeout(() => {
          WS.send(msg);
        }, 1000);
      }

    };

    WS.socket.onmessage = function (event) {
      WS.handler(event.data);
    };

    WS.socket.onclose = function (event) {
      if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
      } else {
        // e.g. server process killed or network down
        // event.code is usually 1006 in this case
        console.log('[close] Connection died');
      }
    };

    WS.socket.onerror = function (error) {
      console.log(`[error] ${error.message}`);
    };
  }
}
