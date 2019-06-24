import {Component} from '@angular/core';
import {UserState} from './model/state/userState';
import {Router} from '@angular/router';
import {WS} from './util/ws';
import {LoadingState} from './model/state/loading';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'py-portal';

  private version = 'v0.42';

  constructor(private router: Router) {
  }

  public isLoading(): boolean {
    return LoadingState.Loading;
  }

  public isLoggedIn(): boolean {
    return UserState.isLoggedIn();
  }

  public logout() {
    UserState.logout();
    const link = ['/'];
    this.router.navigate(link);
  }

  public isGreen(): boolean {
    return WS.connected;
  }

}
