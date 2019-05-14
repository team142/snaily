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
  private theme = 'bg-dark';

  constructor(private router: Router) {
  }

  ngOnInit() {
  }

  public request(): void {
    this.theme = 'bg-dark animated flipOutX';

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
          Swal.fire({
            position: 'middle-end',
            type: 'success',
            title: 'Saved',
            showConfirmButton: false,
            timer: 1800
          });
        }, 1500);

        setTimeout(() => {
          this.router.navigate(['./']);
        }, 3200);

      }
    });
  }
}
