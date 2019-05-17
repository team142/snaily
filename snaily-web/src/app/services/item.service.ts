import {Injectable} from '@angular/core';
import {ItemV1} from '../model/item-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  constructor(private http: HttpClient) {
  }

  public post(item: ItemV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlNewItemV1, item)
      .toPromise()
      .then((result) => {
        win(result);
      })
      .catch((error) => {
        fail(error);
      });

  }

}
