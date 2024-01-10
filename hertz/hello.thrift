namespace go hello.world

service HelloService {
    string Hello(1: string name);
    string Bye(1: string name);
}