import {Component, OnInit} from '@angular/core';
import {ParamMap, Router} from '@angular/router';
import {ItemService} from '../../services/item.service';
import {ActivatedRoute} from '@angular/router';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {


  constructor(private route: ActivatedRoute, private router: Router, private itemService: ItemService) {

  }

  ngOnInit() {

  }

}
