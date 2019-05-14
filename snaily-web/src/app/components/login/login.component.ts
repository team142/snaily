import {Component, OnInit} from '@angular/core';
import {UserState} from '../../model/state/userState';
import {Router} from '@angular/router';
import {RegisterService} from '../../services/register.service';
import {MessageRegisterRequestV1} from '../../model/register-request-v1';
import {MessageLoginRequestV1} from '../../model/login-request-v1';
import {LoginService} from '../../services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  private registerReq: MessageRegisterRequestV1;
  private loginReq: MessageLoginRequestV1;

  constructor(private router: Router, private registerService: RegisterService, private loginService: LoginService) {
    this.registerReq = new MessageRegisterRequestV1();
    this.loginReq = new MessageLoginRequestV1();
  }

  ngOnInit() {
    this.amILoggedIn();
  }

  private amILoggedIn(): void {
    if (UserState.isLoggedIn()) {
      const link = ['./home'];
      this.router.navigate(link);
    }
  }

  public register() {
    this.registerService.post(this.registerReq);
    setTimeout(() => {
      this.amILoggedIn();
    }, 1000);

  }

  public login() {
    this.loginService.post(this.loginReq);
    setTimeout(() => {
      this.amILoggedIn();
    }, 1000);

  }

}
