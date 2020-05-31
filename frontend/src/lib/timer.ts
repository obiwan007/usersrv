let index = 0;
export class TimeEntry {
  id: number = 0;
  description: string = "";
  short: string = "";
  teammember: string = "";
  client: string = "";
  project: string = "";
  tags: string[] = [];
  elapsedSeconds: number = 0;
  timerStart: Date = new Date();
  timerEnd: Date = new Date();
  isRunning: boolean = false;
  get tStart(): string {
    return this.timerStart.toLocaleTimeString();
  }
  get tEnd(): string {
    return this.timerEnd.toLocaleTimeString();
  }
  hms(): string {
    const d = this.elapsed();

    let h = Math.floor(d / 3600);
    let m = Math.floor((d % 3600) / 60);
    let s = Math.floor((d % 3600) % 60);

    return (
      ("0" + h).slice(-2) +
      ":" +
      ("0" + m).slice(-2) +
      ":" +
      ("0" + s).slice(-2)
    );
  }
  /**
   *
   */
  constructor() {
    this.timerStart = new Date();
  }
  elapsed(): number {
    let et = new Date();
    if (!this.isRunning) {
      et = this.timerEnd;
    }
    return (et.getTime() - this.timerStart.getTime()) / 1000;
  }

  static fromRaw(
    entry: any,
    existingEntry: TimeEntry = new TimeEntry()
  ): TimeEntry {
    const o = Object.assign(existingEntry, entry);
    o.timerStart = new Date(o.timerStart);
    o.timerEnd = new Date(o.timerEnd);
    return o;
  }
}

export class Timer {
  private entries: TimeEntry[] = [];
  timerStart: Date = new Date();
  timerEnd: Date = new Date();
  currentTimer: TimeEntry = new TimeEntry();
  constructor() {
    const d = localStorage.getItem("timer");
    if (d) {
      this.entries =
        d && JSON.parse(d).map((entry: any) => TimeEntry.fromRaw(entry));
    }

    const tc = localStorage.getItem("timerCurrent");
    if (tc) {
      this.currentTimer = TimeEntry.fromRaw(JSON.parse(tc));
      console.log("CURRENTTIMER LOADED", this.currentTimer);
    }

    const i = localStorage.getItem("timerIndex");

    index = i ? +i : 0;
  }
  save() {
    console.log("Save:", this.currentTimer);
    localStorage.setItem("timer", JSON.stringify(this.entries));
    localStorage.setItem("timerCurrent", JSON.stringify(this.currentTimer));
    localStorage.setItem("timerIndex", JSON.stringify(index));
  }
  discardTimer() {
    this.currentTimer.timerStart = new Date();
    this.currentTimer.isRunning = false;
    this.currentTimer.description = "";
    this.save();
  }
  getTimer() {
    return this.currentTimer;
  }

  elapsed(): number {
    return this.currentTimer.elapsed();
  }

  static notifyMessage(msg: string, body: string) {
    const permission = Notification.permission;
    console.log("Allowed ", permission)
    if (permission === "granted") {
      const n = new Notification(msg, { body: body, icon: 'logo512.png' });
      n.onclick = () => {
        console.log('Clicked on notification');
      };
    }
  }

  // Entries(filter = "0"): TimeEntry[] {
  //   this.entries.sort(function (a, b) {
  //     if (a.timerStart > b.timerStart) {
  //       return -1;
  //     }
  //     if (b.timerStart > a.timerStart) {
  //       return 1;
  //     }
  //     return 0;
  //   });
  //   console.log('"timefilter', filter);
  //   switch (filter) {
  //     case "0": // show all
  //       return [...this.entries];
  //     case "1": {
  //       // Today
  //       const cd = new Date().toDateString();
  //       const entries = _.filter(this.entries, (e) => {
  //         return e.timerStart.toDateString() === cd;
  //       });
  //       return entries;
  //     }
  //     case "2": {
  //       // Yesterday
  //       const cd = moment().add("day", -1).toDate().toDateString();
  //       console.log(cd);
  //       const entries = _.filter(this.entries, (e) => {
  //         return e.timerStart.toDateString() === cd;
  //       });
  //       return entries;
  //     }
  //     case "7": {
  //       // Start of Week
  //       const cd = moment().startOf("week").toDate();
  //       console.log(cd);
  //       const entries = _.filter(this.entries, (e) => {
  //         return e.timerStart.getTime() > cd.getTime();
  //       });
  //       return entries;
  //     }
  //     case "30": {
  //       // Month
  //       const cd = moment().startOf("month").toDate();
  //       console.log(cd);
  //       const entries = _.filter(this.entries, (e) => {
  //         return e.timerStart.getTime() > cd.getTime();
  //       });
  //       return entries;
  //     }
  //   }
  //   return [...this.entries];
  // }

  static hms(d: number | null | undefined): string {
    if (!d) {
      return "00:00:00";
    }
    let h = Math.floor(d / 3600);
    let m = Math.floor((d % 3600) / 60);
    let s = Math.floor((d % 3600) % 60);

    return (
      ("0" + h).slice(-2) +
      ":" +
      ("0" + m).slice(-2) +
      ":" +
      ("0" + s).slice(-2)
    );
  }
}

export default new Timer();
