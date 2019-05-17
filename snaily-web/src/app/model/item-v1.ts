export class ItemV1 {
  id: string;
  parent: string;
  title: string;
  body: string;
  createdBy: string;
  WaitingFor: string;

  constructor() {
    this.id = '';
    this.parent = '';
    this.title = '';
    this.body = '';
    this.createdBy = '';
    this.WaitingFor = '';
  }

}
