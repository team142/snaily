export class ItemV1 {
  id: string;
  parent: string;
  title: string;
  body: string;
  createdBy: string;
  waitingFor: string;
  waitingForDone: boolean;
  createdByDone: boolean;

  constructor() {
    this.id = '';
    this.parent = '';
    this.title = '';
    this.body = '';
    this.createdBy = '';
    this.waitingFor = '';
    this.waitingForDone = false;
    this.createdByDone = false;

  }

}
