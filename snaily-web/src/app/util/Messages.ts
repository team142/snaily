declare var Swal: any;

export class Messages {
  static AccessDenied() {
    Swal.fire({
      position: 'middle-end',
      type: 'error',
      title: 'Access Denied',
      showConfirmButton: false,
      timer: 1800
    });

  }
}
