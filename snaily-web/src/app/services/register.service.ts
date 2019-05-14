import {Injectable} from '@angular/core';
import {MessageRegisterRequestV1} from '../model/register-request-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private http: HttpClient) {
  }

  public post(item: MessageRegisterRequestV1): void {
    this.http.post(environment.urlRegisterV1, item)
      .toPromise()
      .then((result => this.handleRegiserResult(result)))
      .catch((error) => console.log(error));

  }

  private handleRegiserResult(result: any) {
    if (result.ok === true) {
      alert('You are registered. Try logging in next.');
    } else {
      alert('Failed to register. Perhaps you are already registered?');
    }

  }

}
