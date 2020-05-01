export class TimeEntry {
  teammember: string = "";
  client: string = "";
  project: string = "";
  tags: string[] = [];
  elapsedSeconds: number = 0;
  timerStart: Date = new Date();
  timerEnd: Date = new Date();
  description: string = "";
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
}

export class Timer {
  entries: TimeEntry[] = [];
  timerStart: Date = new Date();
  timerEnd: Date = new Date();
  currentTimer: TimeEntry = new TimeEntry();
  startTimer() {
    this.currentTimer = new TimeEntry();
    this.currentTimer.isRunning = true;
  }
  getTimer() {
    return this.currentTimer;
  }
  endTimer() {
    this.currentTimer.timerEnd = new Date();
    this.currentTimer.isRunning = false;
    this.currentTimer.elapsedSeconds = this.currentTimer.elapsed();
    this.entries.push(this.currentTimer);
  }
  elapsed(): number {
    return this.currentTimer.elapsed();
  }
}

export default new Timer();
