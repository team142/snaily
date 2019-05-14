import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

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

  constructor(private router: Router) {
  }

  ngOnInit() {
    this.setTimeAsNow();
  }

  public request(): void {
    this.theme = 'animated flipOutX';

    setTimeout(() => {
        this.save();
      }
      , 1500);
  }

  private save(): void {
    let timerInterval;
    Swal.fire({
      title: 'Saving',
      html: 'Checking in <strong></strong> ms.',
      timer: 3000,
      onBeforeOpen: () => {
        Swal.showLoading();
        timerInterval = setInterval(() => {
          Swal.getContent().querySelector('strong')
            .textContent = Swal.getTimerLeft();
        }, 100);
      },
      onClose: () => {
        clearInterval(timerInterval);
      }
    }).then((result) => {
      if (
        // Read more about handling dismissals
        result.dismiss === Swal.DismissReason.timer
      ) {
        this.theme = 'bg-success animated flipInX';
        setTimeout(() => {
          this.savedSuccess();
        }, 1500);

        setTimeout(() => {
          this.router.navigate(['./']);
        }, 3200);

      }
    });
  }

  private savedSuccess(): void {
    Swal.fire({
      position: 'middle-end',
      type: 'success',
      title: 'Saved',
      showConfirmButton: false,
      timer: 1800
    });

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
