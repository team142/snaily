import {Injectable} from '@angular/core';
import {ItemV1} from '../model/item-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {MyItemsRequestV1} from '../model/my-items-request-v1';

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

  public getMyItems(item: MyItemsRequestV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlMyItemsV1, item)
      .toPromise()
      .then((result) => {
        win(result);
      })
      .catch((error) => {
        fail(error);
      });
  }

  public getItem(item: ItemV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlGetItemV1, item)
      .toPromise()
      .then((result) => {
        win(result);
      })
      .catch((error) => {
        fail(error);
      });
  }


}
