import {Injectable} from '@angular/core';
import {MessageLoginRequestV1} from '../model/login-request-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {UserState} from '../model/state/userState';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) {
  }

  public post(item: MessageLoginRequestV1): void {
    this.http.post(environment.urlLoginV1, item)
      .toPromise()
      .then((result => this.handleResult(result)))
      .catch((error) => console.log(error));
  }

  private handleResult(result: any) {
    if (result.ok === true) {
      UserState.login();
      UserState.setKey(result.key);
      console.log(result.key);
    } else {
      alert('Authentication failed');
    }
  }


}
