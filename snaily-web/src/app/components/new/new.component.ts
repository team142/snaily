import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {ItemService} from '../../services/item.service';
import {ItemV1} from '../../model/item-v1';
import {UserState} from '../../model/state/userState';

declare var Swal: any;

@Component({
  selector: 'app-maintain',
  templateUrl: './new.component.html',
  styleUrls: ['./new.component.css']
})
export class NewComponent implements OnInit {

  private email: string;
  private body: string;
  private title: string;
  private day: number;
  private nmonth: number;
  private year: number;
  private hour: number;
  private minute: number;


  private theme = '';

  constructor(private router: Router, private itemService: ItemService) {
  }

  ngOnInit() {
    this.setTimeAsNow();
  }

  public request(): void {
    this.saveItem();
  }


  private saveItem(): void {

    const i = new ItemV1();
    i.title = this.title;
    i.body = this.body;
    i.createdBy = UserState.getMyID();
    i.WaitingFor = this.email;


    this.itemService.post(i, (result) => (
        this.animateSuccess(result)
      ),
      (error) => {
        alert('Error - ' + error);
      }
    );
  }

  private animateSuccess(result: any) {
    console.log(result);
    Swal.fire({
      position: 'middle-end',
      type: 'success',
      title: 'Saved',
      showConfirmButton: false,
      timer: 1800
    });

    setTimeout(() => {
      this.router.navigate(['./']);
    }, 3200);

  }


  private setTimeAsNow() {

    const n = new Date();

    this.nmonth = n.getUTCMonth();
    this.year = n.getFullYear();

    this.day = n.getUTCDate();
    this.hour = n.getUTCHours();
    this.minute = n.getUTCMinutes();

  }
}
