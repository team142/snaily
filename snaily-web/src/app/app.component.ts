import {Component} from '@angular/core';
import {UserState} from './model/state/userState';
import {Router} from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'py-portal';


  constructor(private router: Router) {
  }

  public isLoggedIn(): boolean {
    return UserState.isLoggedIn();
  }

  public logout() {
    UserState.logout();
    const link = ['/'];
    this.router.navigate(link);
  }


}
