import {Component, OnInit} from '@angular/core';
import {UserState} from '../../model/state/userState';
import {Router} from '@angular/router';
import {RegisterService} from '../../services/register.service';
import {MessageRegisterRequestV1} from '../../model/register-request-v1';
import {MessageLoginRequestV1} from '../../model/login-request-v1';
import {LoginService} from '../../services/login.service';

declare var Swal: any;

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  private registerReq: MessageRegisterRequestV1;
  private loginReq: MessageLoginRequestV1;

  private themeRegister = 'animated bounceInRight';
  private themeLogin = 'animated bounceInLeft';

  constructor(private router: Router, private registerService: RegisterService, private loginService: LoginService) {
    this.registerReq = new MessageRegisterRequestV1();
    this.loginReq = new MessageLoginRequestV1();
  }

  ngOnInit() {
    this.amILoggedIn();
  }

  private amILoggedIn(): void {
    if (UserState.isLoggedIn()) {
      this.router.navigate(['./home']);
    }
  }

  public register() {
    this.themeRegister = 'animated bounceOutRight';

    this.registerService.post(this.registerReq, (result) => {
      alert('Ok. Next you must login');
    }, (error) => {
      alert('Could not register? Maybe you\'re already a user?');
    });

  }

  public login() {
    this.flyOut();


    this.loginService.post(this.loginReq, (result) => {
      if (result.ok === true) {
        UserState.login();
        UserState.setKey(result.key);

        Swal.fire({
          position: 'middle-end',
          type: 'success',
          title: 'Welcome',
          showConfirmButton: false,
          timer: 1300
        });


      } else {
        setTimeout(() => {
          this.flyIn();
        }, 1000);

        Swal.fire({
          position: 'middle-end',
          type: 'error',
          title: 'Incorrect email or password',
          showConfirmButton: true,
        });
      }

    }, (error) => {
      alert(error);
    });


    setTimeout(() => {
      this.amILoggedIn();
    }, 1000);

  }

  private flyOut() {
    this.themeRegister = 'animated bounceOutRight';
    this.themeLogin = 'animated bounceOutLeft';

  }

  private flyIn() {
    this.themeRegister = 'animated bounceInRight';
    this.themeLogin = 'animated bounceInLeft';

  }

}
