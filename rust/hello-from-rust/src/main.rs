use fake::faker::company::raw::*;
use fake::faker::name::raw::*;
use fake::locales::*;
use fake::Fake;

fn main() {
    let name: String = Name(EN).fake();
    let company: String = CompanyName(EN).fake();
    let buzz1: String = Buzzword(EN).fake();
    let buzz2: String = BuzzwordMiddle(EN).fake();
    let buzz3: String = BuzzwordTail(EN).fake();
    println!(
        "Hi! My name is \"{name}\", and I work in \"{company}\" doing \"{buzz1} {buzz2} {buzz3}\""
    );
}
