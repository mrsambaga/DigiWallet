export const SetCookie = (
  cname: string,
  cvalue: string,
  exdays: number,
): void => {
  const d = new Date();
  d.setTime(d.getTime() + exdays * 24 * 60 * 60 * 1000);
  const expires = 'expires=' + d.toUTCString();
  document.cookie = cname + '=' + cvalue + ';' + expires + ';path=/';
};

export const GetCookie = (cname: string): string => {
  const name = cname + '=';
  const ca = document.cookie.split(';');
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return '';
};

// export const RefreshCookie = (cname: string): void => {
//   const cookie = GetCookie(cname);
//   if (cookie) {
//     SetCookie("")
//   }
// };

export const CheckCookie = (): void => {
  let user = GetCookie('username');
  if (user != '') {
    alert('Welcome again ' + user);
  } else {
    const userPrompted = prompt('Please enter your name:', '')!;
    if (userPrompted != '' && userPrompted != null) {
      user = userPrompted;
      SetCookie('username', user, 365);
    }
  }
};
