import {Component, OnInit} from '@angular/core';
import {ParamMap, Router} from '@angular/router';
import {ItemService} from '../../services/item.service';
import {ActivatedRoute} from '@angular/router';
import {HttpErrorResponse} from '@angular/common/http';
import {Messages} from '../../util/Messages';
import {ItemV1} from '../../model/item-v1';
import {UserV1} from '../../model/user-v1';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  private createdByMe: ItemV1[] = [];
  private waitingForMe: ItemV1[] = [];
  private users: UserV1[] = [];

  constructor(private route: ActivatedRoute, private router: Router, private itemService: ItemService) {

  }

  ngOnInit() {
    this.refresh();
  }


  private refresh() {
    this.itemService
      .getMyItems((result) => {
          this.users = result.users;
          this.createdByMe = result.createdByMe;
          this.waitingForMe = result.waitingForMe;
          console.log(result);


        }, (err: HttpErrorResponse) => {
          if (err.status === 403) {
            Messages.AccessDenied();
            this.router.navigate(['./login']);
            return;
          }
        }
      );
  }

  public getUserFullName(id: string): string {
    for (const o of this.users) {
      if (o.id === id) {
        if (o.firstName === '' || o.lastName === '') {
          return o.email;
        }
        const r = o.firstName + ' ' + o.lastName;
        return r;
      }
    }
    return '?';
  }

}
