import {Injectable} from '@angular/core';
import {MessageLoginRequestV1} from '../model/login-request-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) {
  }

  public post(item: MessageLoginRequestV1): void {
    this.http.post(environment.urlLoginV1, item)
      .toPromise()
      .then((result => this.handleRegiserResult(result)))
      .catch((error) => console.log(error));

  }

  private handleRegiserResult(result: any) {
    console.log(result);
  }


}
