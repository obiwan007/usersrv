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
    const i = localStorage.getItem("timerIndex");

    index = i ? +i : 0;
  }
  save() {
    localStorage.setItem("timer", JSON.stringify(this.entries));
    localStorage.setItem("timerIndex", JSON.stringify(index));
  }
  startTimer() {
    this.currentTimer = new TimeEntry();
    index = index + 1;
    this.currentTimer.id = index;
    this.currentTimer.isRunning = true;
    this.save();
  }
  getTimer() {
    return this.currentTimer;
  }
  endTimer() {
    this.currentTimer.timerEnd = new Date();
    this.currentTimer.isRunning = false;
    this.currentTimer.elapsedSeconds = this.currentTimer.elapsed();
    this.entries.push(this.currentTimer);
    this.save();
  }
  elapsed(): number {
    return this.currentTimer.elapsed();
  }
  del(p: TimeEntry) {
    let data = this.entries;
    const entry = data.find((d) => d.id === p.id);

    if (entry) {
      console.log("Delete ", p);
      const index = data.indexOf(entry);
      data.splice(index, 1);
    }
    this.save();
    return this.Entries();
  }

  update(p: TimeEntry) {
    let data = this.entries;
    const entry = data.find((d) => d.id === p.id);

    if (entry) {
      console.log("Delete ", p);
      TimeEntry.fromRaw(p, entry);
    }
    this.save();
    return this.Entries();
  }
  Entries(): TimeEntry[] {
    return [...this.entries];
  }
}

export default new Timer();
