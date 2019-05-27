import {Component, Input, OnInit} from '@angular/core';
import {ActivatedRoute, Params, Router} from '@angular/router';
import {ItemService} from '../../services/item.service';
import {ItemV1} from '../../model/item-v1';
import {HttpErrorResponse} from '@angular/common/http';
import {Messages} from '../../util/Messages';

declare var Swal: any;

@Component({
  selector: 'app-browse',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css']
})
export class ViewComponent implements OnInit {
  @Input() id: string;

  constructor(
    private route: ActivatedRoute,
    private itemService: ItemService,
    private router: Router,
  ) {
  }

  ngOnInit() {

    // Heresy
    this.id = window.location.search.split('=')[1];

    this.route.params.forEach((params: Params) => {
      if (params['id'] !== undefined) {
        this.id = params['id'];
      }
    });

    this.load();

  }

  private load(): void {
    const i = new ItemV1();
    i.id = this.id;
    this.itemService.getItem(i, (result) => {
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
}
