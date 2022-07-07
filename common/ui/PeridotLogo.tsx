/*
 * Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
 * Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
 * Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 * this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation
 * and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its contributors
 * may be used to endorse or promote products derived from this software without
 * specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

import React from 'react';

export interface PeridotLogoProps {
  className?: string;
}

export const PeridotLogo = (props: PeridotLogoProps) => {
  return (
    <svg
      height="87"
      viewBox="0 0 551 87"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={props.className}
    >
      <path
        d="M15.4375 42.5938H18.5625C21.625 42.5938 23.8854 42.0833 25.3438 41.0625C26.8021 40.0417 27.5312 38.4375 27.5312 36.25C27.5312 34.0833 26.7812 32.5417 25.2812 31.625C23.8021 30.7083 21.5 30.25 18.375 30.25H15.4375V42.5938ZM15.4375 50.4688V68H5.75V22.3125H19.0625C25.2708 22.3125 29.8646 23.4479 32.8438 25.7188C35.8229 27.9688 37.3125 31.3958 37.3125 36C37.3125 38.6875 36.5729 41.0833 35.0938 43.1875C33.6146 45.2708 31.5208 46.9062 28.8125 48.0938C35.6875 58.3646 40.1667 65 42.25 68H31.5L20.5938 50.4688H15.4375ZM74.3125 68H48V22.3125H74.3125V30.25H57.6875V40.2812H73.1562V48.2188H57.6875V60H74.3125V68ZM110.844 55.3125C110.844 59.4375 109.354 62.6875 106.375 65.0625C103.417 67.4375 99.2917 68.625 94 68.625C89.125 68.625 84.8125 67.7083 81.0625 65.875V56.875C84.1458 58.25 86.75 59.2188 88.875 59.7812C91.0208 60.3438 92.9792 60.625 94.75 60.625C96.875 60.625 98.5 60.2188 99.625 59.4062C100.771 58.5938 101.344 57.3854 101.344 55.7812C101.344 54.8854 101.094 54.0938 100.594 53.4062C100.094 52.6979 99.3542 52.0208 98.375 51.375C97.4167 50.7292 95.4479 49.6979 92.4688 48.2812C89.6771 46.9688 87.5833 45.7083 86.1875 44.5C84.7917 43.2917 83.6771 41.8854 82.8438 40.2812C82.0104 38.6771 81.5938 36.8021 81.5938 34.6562C81.5938 30.6146 82.9583 27.4375 85.6875 25.125C88.4375 22.8125 92.2292 21.6562 97.0625 21.6562C99.4375 21.6562 101.698 21.9375 103.844 22.5C106.01 23.0625 108.271 23.8542 110.625 24.875L107.5 32.4062C105.062 31.4062 103.042 30.7083 101.438 30.3125C99.8542 29.9167 98.2917 29.7188 96.75 29.7188C94.9167 29.7188 93.5104 30.1458 92.5312 31C91.5521 31.8542 91.0625 32.9688 91.0625 34.3438C91.0625 35.1979 91.2604 35.9479 91.6562 36.5938C92.0521 37.2188 92.6771 37.8333 93.5312 38.4375C94.4062 39.0208 96.4583 40.0833 99.6875 41.625C103.958 43.6667 106.885 45.7188 108.469 47.7812C110.052 49.8229 110.844 52.3333 110.844 55.3125ZM128.656 68H119.125V22.3125H145.312V30.25H128.656V42.0312H144.156V49.9375H128.656V68ZM171.406 22.3125H184.312C190.375 22.3125 194.76 23.2188 197.469 25.0312C200.177 26.8438 201.531 29.7083 201.531 33.625C201.531 36.3333 200.771 38.5729 199.25 40.3438C197.75 42.0938 195.552 43.2292 192.656 43.75V44.0625C199.594 45.25 203.062 48.8958 203.062 55C203.062 59.0833 201.677 62.2708 198.906 64.5625C196.156 66.8542 192.302 68 187.344 68H171.406V22.3125ZM176.719 41.875H185.469C189.219 41.875 191.917 41.2917 193.562 40.125C195.208 38.9375 196.031 36.9479 196.031 34.1562C196.031 31.5938 195.115 29.75 193.281 28.625C191.448 27.4792 188.531 26.9062 184.531 26.9062H176.719V41.875ZM176.719 46.375V63.4688H186.25C189.938 63.4688 192.708 62.7604 194.562 61.3438C196.438 59.9062 197.375 57.6667 197.375 54.625C197.375 51.7917 196.417 49.7083 194.5 48.375C192.604 47.0417 189.708 46.375 185.812 46.375H176.719ZM217 33.75V55.9688C217 58.7604 217.635 60.8438 218.906 62.2188C220.177 63.5938 222.167 64.2812 224.875 64.2812C228.458 64.2812 231.073 63.3021 232.719 61.3438C234.385 59.3854 235.219 56.1875 235.219 51.75V33.75H240.406V68H236.125L235.375 63.4062H235.094C234.031 65.0938 232.552 66.3854 230.656 67.2812C228.781 68.1771 226.635 68.625 224.219 68.625C220.052 68.625 216.927 67.6354 214.844 65.6562C212.781 63.6771 211.75 60.5104 211.75 56.1562V33.75H217ZM256.625 68H251.438V33.75H256.625V68ZM251 24.4688C251 23.2812 251.292 22.4167 251.875 21.875C252.458 21.3125 253.188 21.0312 254.062 21.0312C254.896 21.0312 255.615 21.3125 256.219 21.875C256.823 22.4375 257.125 23.3021 257.125 24.4688C257.125 25.6354 256.823 26.5104 256.219 27.0938C255.615 27.6562 254.896 27.9375 254.062 27.9375C253.188 27.9375 252.458 27.6562 251.875 27.0938C251.292 26.5104 251 25.6354 251 24.4688ZM272.812 68H267.625V19.375H272.812V68ZM307.125 63.4062H306.844C304.448 66.8854 300.865 68.625 296.094 68.625C291.615 68.625 288.125 67.0938 285.625 64.0312C283.146 60.9688 281.906 56.6146 281.906 50.9688C281.906 45.3229 283.156 40.9375 285.656 37.8125C288.156 34.6875 291.635 33.125 296.094 33.125C300.74 33.125 304.302 34.8125 306.781 38.1875H307.188L306.969 35.7188L306.844 33.3125V19.375H312.031V68H307.812L307.125 63.4062ZM296.75 64.2812C300.292 64.2812 302.854 63.3229 304.438 61.4062C306.042 59.4688 306.844 56.3542 306.844 52.0625V50.9688C306.844 46.1146 306.031 42.6562 304.406 40.5938C302.802 38.5104 300.229 37.4688 296.688 37.4688C293.646 37.4688 291.312 38.6562 289.688 41.0312C288.083 43.3854 287.281 46.7188 287.281 51.0312C287.281 55.4062 288.083 58.7083 289.688 60.9375C291.292 63.1667 293.646 64.2812 296.75 64.2812ZM366.25 55.8438C366.25 59.8646 364.792 63 361.875 65.25C358.958 67.5 355 68.625 350 68.625C344.583 68.625 340.417 67.9271 337.5 66.5312V61.4062C339.375 62.1979 341.417 62.8229 343.625 63.2812C345.833 63.7396 348.021 63.9688 350.188 63.9688C353.729 63.9688 356.396 63.3021 358.188 61.9688C359.979 60.6146 360.875 58.7396 360.875 56.3438C360.875 54.7604 360.552 53.4688 359.906 52.4688C359.281 51.4479 358.219 50.5104 356.719 49.6562C355.24 48.8021 352.979 47.8333 349.938 46.75C345.688 45.2292 342.646 43.4271 340.812 41.3438C339 39.2604 338.094 36.5417 338.094 33.1875C338.094 29.6667 339.417 26.8646 342.062 24.7812C344.708 22.6979 348.208 21.6562 352.562 21.6562C357.104 21.6562 361.281 22.4896 365.094 24.1562L363.438 28.7812C359.667 27.1979 356 26.4062 352.438 26.4062C349.625 26.4062 347.427 27.0104 345.844 28.2188C344.26 29.4271 343.469 31.1042 343.469 33.25C343.469 34.8333 343.76 36.1354 344.344 37.1562C344.927 38.1562 345.906 39.0833 347.281 39.9375C348.677 40.7708 350.802 41.6979 353.656 42.7188C358.448 44.4271 361.74 46.2604 363.531 48.2188C365.344 50.1771 366.25 52.7188 366.25 55.8438ZM369.375 33.75H374.938L382.438 53.2812C384.083 57.7396 385.104 60.9583 385.5 62.9375H385.75C386.021 61.875 386.583 60.0625 387.438 57.5C388.312 54.9167 391.146 47 395.938 33.75H401.5L386.781 72.75C385.323 76.6042 383.615 79.3333 381.656 80.9375C379.719 82.5625 377.333 83.375 374.5 83.375C372.917 83.375 371.354 83.1979 369.812 82.8438V78.6875C370.958 78.9375 372.24 79.0625 373.656 79.0625C377.219 79.0625 379.76 77.0625 381.281 73.0625L383.188 68.1875L369.375 33.75ZM429.156 58.6562C429.156 61.8438 427.969 64.3021 425.594 66.0312C423.219 67.7604 419.885 68.625 415.594 68.625C411.052 68.625 407.51 67.9062 404.969 66.4688V61.6562C406.615 62.4896 408.375 63.1458 410.25 63.625C412.146 64.1042 413.969 64.3438 415.719 64.3438C418.427 64.3438 420.51 63.9167 421.969 63.0625C423.427 62.1875 424.156 60.8646 424.156 59.0938C424.156 57.7604 423.573 56.625 422.406 55.6875C421.26 54.7292 419.01 53.6042 415.656 52.3125C412.469 51.125 410.198 50.0938 408.844 49.2188C407.51 48.3229 406.51 47.3125 405.844 46.1875C405.198 45.0625 404.875 43.7188 404.875 42.1562C404.875 39.3646 406.01 37.1667 408.281 35.5625C410.552 33.9375 413.667 33.125 417.625 33.125C421.312 33.125 424.917 33.875 428.438 35.375L426.594 39.5938C423.156 38.1771 420.042 37.4688 417.25 37.4688C414.792 37.4688 412.938 37.8542 411.688 38.625C410.438 39.3958 409.812 40.4583 409.812 41.8125C409.812 42.7292 410.042 43.5104 410.5 44.1562C410.979 44.8021 411.74 45.4167 412.781 46C413.823 46.5833 415.823 47.4271 418.781 48.5312C422.844 50.0104 425.583 51.5 427 53C428.438 54.5 429.156 56.3854 429.156 58.6562ZM448.688 64.3438C449.604 64.3438 450.49 64.2812 451.344 64.1562C452.198 64.0104 452.875 63.8646 453.375 63.7188V67.6875C452.812 67.9583 451.979 68.1771 450.875 68.3438C449.792 68.5312 448.812 68.625 447.938 68.625C441.312 68.625 438 65.1354 438 58.1562V37.7812H433.094V35.2812L438 33.125L440.188 25.8125H443.188V33.75H453.125V37.7812H443.188V57.9375C443.188 60 443.677 61.5833 444.656 62.6875C445.635 63.7917 446.979 64.3438 448.688 64.3438ZM474.719 68.625C469.656 68.625 465.656 67.0833 462.719 64C459.802 60.9167 458.344 56.6354 458.344 51.1562C458.344 45.6354 459.698 41.25 462.406 38C465.135 34.75 468.792 33.125 473.375 33.125C477.667 33.125 481.062 34.5417 483.562 37.375C486.062 40.1875 487.312 43.9062 487.312 48.5312V51.8125H463.719C463.823 55.8333 464.833 58.8854 466.75 60.9688C468.688 63.0521 471.406 64.0938 474.906 64.0938C478.594 64.0938 482.24 63.3229 485.844 61.7812V66.4062C484.01 67.1979 482.271 67.7604 480.625 68.0938C479 68.4479 477.031 68.625 474.719 68.625ZM473.312 37.4688C470.562 37.4688 468.365 38.3646 466.719 40.1562C465.094 41.9479 464.135 44.4271 463.844 47.5938H481.75C481.75 44.3229 481.021 41.8229 479.562 40.0938C478.104 38.3438 476.021 37.4688 473.312 37.4688ZM539.844 68V45.7188C539.844 42.9896 539.26 40.9479 538.094 39.5938C536.927 38.2188 535.115 37.5312 532.656 37.5312C529.427 37.5312 527.042 38.4583 525.5 40.3125C523.958 42.1667 523.188 45.0208 523.188 48.875V68H518V45.7188C518 42.9896 517.417 40.9479 516.25 39.5938C515.083 38.2188 513.26 37.5312 510.781 37.5312C507.531 37.5312 505.146 38.5104 503.625 40.4688C502.125 42.4062 501.375 45.5938 501.375 50.0312V68H496.188V33.75H500.406L501.25 38.4375H501.5C502.479 36.7708 503.854 35.4688 505.625 34.5312C507.417 33.5938 509.417 33.125 511.625 33.125C516.979 33.125 520.479 35.0625 522.125 38.9375H522.375C523.396 37.1458 524.875 35.7292 526.812 34.6875C528.75 33.6458 530.958 33.125 533.438 33.125C537.312 33.125 540.208 34.125 542.125 36.125C544.062 38.1042 545.031 41.2812 545.031 45.6562V68H539.844Z"
        fill="white"
      />
    </svg>
  );
};
